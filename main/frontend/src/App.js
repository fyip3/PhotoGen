import React, { useState } from 'react';
import ImageUploader from './components/ImageUploader';
import ImageGrid from './components/ImageGrid';
import './App.css';

function App() {
    const [images, setImages] = useState([]);

    return (
        <div className="App">
            <h1>LinkedIn Profile Picture Generator</h1>
            <div className="content-wrapper">
                <ImageUploader setImages={setImages} />
                <ImageGrid images={images} />
            </div>
        </div>
    );
}

export default App;

