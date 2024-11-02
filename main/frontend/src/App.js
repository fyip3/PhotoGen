import React, { useState } from 'react';
import ImageUploader from './components/ImageUploader';
import ImageGrid from './components/ImageGrid';

function App() {
    const [images, setImages] = useState([]);

    return (
        <div className="App">
            <h1>LinkedIn Profile Picture Generator</h1>
            <ImageUploader setImages={setImages} />
            <ImageGrid images={images} />
        </div>
    );
}

export default App;
