import React, { useState, useEffect, ReactNode } from 'react';
import styles from './styles.module.css';

interface BookReaderProps {
    children: ReactNode;
}

export default function BookReader({ children }: BookReaderProps) {
    const [readingProgress, setReadingProgress] = useState(0);
    const [showControls, setShowControls] = useState(true);
    const [fontSize, setFontSize] = useState(16);

    // Calculate reading progress
    useEffect(() => {
        const handleScroll = () => {
            const totalHeight = document.body.scrollHeight - window.innerHeight;
            const progress = (window.scrollY / totalHeight) * 100;
            setReadingProgress(progress);
        };

        window.addEventListener('scroll', handleScroll);
        return () => window.removeEventListener('scroll', handleScroll);
    }, []);

    // Toggle reader controls
    const toggleControls = () => {
        setShowControls(!showControls);
    };

    // Font size controls
    const increaseFontSize = () => {
        setFontSize(prevSize => Math.min(prevSize + 1, 24));
        document.documentElement.style.setProperty('--reader-font-size', `${fontSize + 1}px`);
    };

    const decreaseFontSize = () => {
        setFontSize(prevSize => Math.max(prevSize - 1, 12));
        document.documentElement.style.setProperty('--reader-font-size', `${fontSize - 1}px`);
    };

    return (
        <div className={styles.readerContainer}>
            <div className={styles.progressBar} style={{ width: `${readingProgress}%` }}></div>

            {showControls && (
                <div className={styles.readerControls}>
                    <div className={styles.fontControls}>
                        <button onClick={decreaseFontSize} aria-label="Decrease font size">A-</button>
                        <span>{fontSize}px</span>
                        <button onClick={increaseFontSize} aria-label="Increase font size">A+</button>
                    </div>

                    <button
                        className={styles.toggleTheme}
                        onClick={() => {
                            document.documentElement.classList.toggle('dark');
                        }}
                        aria-label="Toggle dark mode"
                    >
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path d="M12 16C14.2091 16 16 14.2091 16 12C16 9.79086 14.2091 8 12 8C9.79086 8 8 9.79086 8 12C8 14.2091 9.79086 16 12 16Z" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                            <path d="M12 2V4" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                            <path d="M12 20V22" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                            <path d="M4.92993 4.93005L6.33993 6.34005" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                            <path d="M17.6599 17.66L19.0699 19.07" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                            <path d="M2 12H4" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                            <path d="M20 12H22" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                            <path d="M6.33993 17.66L4.92993 19.07" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                            <path d="M19.0699 4.93005L17.6599 6.34005" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
                        </svg>
                    </button>
                </div>
            )}

            <div
                className={styles.readerContent}
                style={{
                    '--reader-font-size': `${fontSize}px`,
                } as React.CSSProperties}
            >
                {children}
            </div>

            <button
                className={styles.toggleControls}
                onClick={toggleControls}
                aria-label={showControls ? "Hide controls" : "Show controls"}
            >
                {showControls ? '−' : '+'}
            </button>
        </div>
    );
} 