import React, { useState } from 'react';
import axios from 'axios';

function ImageUploader({ setImages }) {
    const [file, setFile] = useState(null);
    const [description, setDescription] = useState('');
    const [loading, setLoading] = useState(false);

    const handleUpload = async () => {
        if (!file) {
            alert("Please select a file first!");
            return;
        }

        const allowedTypes = ["image/jpeg", "image/png"];
        if (!allowedTypes.includes(file.type)) {
            alert("Unsupported file type. Please upload a JPEG or PNG image.");
            return;
        }

        const formData = new FormData();
        formData.append('file', file);
        formData.append('description', description); 
        setLoading(true);
        try {
            const response = await axios.post("http://localhost:8080/upload", formData, {
                headers: { "Content-Type": "multipart/form-data" },
            });
            setImages(response.data.generated_images);
            console.log("Generated Images:", response.data.generated_images); 

        } catch (error) {
            console.error("Error uploading file:", error);
            alert("Failed to upload the file.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="uploader-wrapper">
            <h2>Upload Your Image</h2>
            <input type="file" onChange={(e) => setFile(e.target.files[0])} />
            <br />
            <input
                type="text"
                placeholder="Enter description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
            />
            <br />
            <button onClick={handleUpload} disabled={loading}>
                {loading ? "Generating..." : "Generate Images"}
            </button>
            {loading && <div className="loader"></div>} {/* Loader */}
        </div>
    );
}

export default ImageUploader;
