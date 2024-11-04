import { createTheme } from '@mui/material/styles';

const theme = createTheme({
    palette: {
        mode: 'dark',  
        primary: {
            main: '#4caf50',  
        },
        secondary: {
            main: '#3e8e41',
        },
        background: {
            default: '#1e1e2f',  
            paper: '#28293d',   
        },
        text: {
            primary: '#ffffff',
        },
    },
    typography: {
        fontFamily: 'Arial, sans-serif',
    },
    components: {
        MuiButton: {
            styleOverrides: {
                root: {
                    textTransform: 'none',  
                    color: '#ffffff',
                },
                containedPrimary: {
                    backgroundColor: '#4caf50',
                    '&:hover': {
                        backgroundColor: '#45a049',
                    },
                },
            },
        },
    },
});

export default theme;
