import React from 'react';

function ImageGrid({ images }) {
    return (
        <div className="image-grid">
            {images.length > 0 ? (
                images.map((img, index) => (
                    <div key={index} className="image-item">
                        <img src={img.url} alt={`Generated option ${index + 1}`} />
                    </div>
                ))
            ) : (
                <p>No images generated yet. Please upload an image to start.</p>
            )}
        </div>
    );
}

export default ImageGrid;
