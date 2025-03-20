import React, { useState, useEffect } from 'react';
import DocItem from '@theme-original/DocItem';
import BookReader from '@site/src/components/BookReader';
import Admonition from '@theme/Admonition';
import { useLocation } from '@docusaurus/router';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';

export default function DocItemWrapper(props) {
  const location = useLocation();
  const { siteConfig } = useDocusaurusContext();
  const [readingTime, setReadingTime] = useState(null);
  const [wordCount, setWordCount] = useState(null);

  useEffect(() => {
    // Calculate reading time after component mounts and content is available
    setTimeout(() => {
      const article = document.querySelector('article.markdown');
      if (article) {
        const text = article.textContent || '';
        const words = text.split(/\s+/).filter(Boolean).length;
        const wpm = 200; // Average reading speed (words per minute)
        const time = Math.ceil(words / wpm);
        
        setWordCount(words);
        setReadingTime(time);
      }
    }, 100);
  }, [location.pathname]); // Recalculate when path changes

  // Determine if this is a book page
  const isBookPage = location.pathname.includes('/docs/');

  return (
    <>
      {isBookPage && readingTime && (
        <div style={{ 
          margin: '1rem 0 2rem', 
          padding: '0.75rem 1rem', 
          backgroundColor: 'var(--ifm-color-emphasis-100)',
          borderRadius: '8px',
          fontSize: '0.9rem',
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center'
        }}>
          <div>
            <strong>Reading time:</strong> {readingTime} minute{readingTime !== 1 ? 's' : ''}
            <span style={{ margin: '0 0.5rem' }}>•</span>
            <strong>Words:</strong> {wordCount}
          </div>
          <button 
            onClick={() => document.body.classList.toggle('immersive-reading')}
            style={{
              background: 'transparent',
              border: 'none',
              cursor: 'pointer',
              display: 'flex',
              alignItems: 'center',
              color: 'var(--ifm-color-primary)',
            }}
          >
            Toggle Immersive Mode
          </button>
        </div>
      )}
      
      {isBookPage ? (
        <BookReader>
          <DocItem {...props} />
        </BookReader>
      ) : (
        <DocItem {...props} />
      )}
      
      {isBookPage && (
        <div className="book-navigation">
          <a href="#" className="prev-link">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M19 12H5M12 19l-7-7 7-7"/>
            </svg>
            Previous Chapter
          </a>
          <a href="#" className="next-link">
            Next Chapter
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M5 12h14M12 5l7 7-7 7"/>
            </svg>
          </a>
        </div>
      )}
    </>
  );
} 