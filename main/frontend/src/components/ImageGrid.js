import React from 'react';
import { Grid, Card, CardMedia, Typography, Skeleton, Fade } from '@mui/material';

function ImageGrid({ images }) {
    return (
        <Grid container spacing={2} justifyContent="center" sx={{ marginTop: '20px' }}>
            {images.length > 0 ? (
                images.map((img, index) => (
                    <Fade in={true} timeout={800} key={index}>
                        <Grid item xs={12} sm={6} md={4} lg={3}>
                            <Card>
                                <CardMedia
                                    component="img"
                                    height="200"
                                    image={img.url}
                                    alt={`Generated option ${index + 1}`}
                                    loading="lazy"
                                    sx={{ objectFit: 'cover', borderRadius: '8px' }}
                                />
                                <Typography variant="body2" align="center" sx={{ padding: '10px' }}>
                                    Generated option {index + 1}
                                </Typography>
                            </Card>
                        </Grid>
                    </Fade>
                ))
            ) : (
                <Typography variant="body2" align="center" sx={{ marginTop: '20px' }}>
                    No images generated yet. Please upload an image to start.
                </Typography>
            )}
        </Grid>
    );
}

export default ImageGrid;
