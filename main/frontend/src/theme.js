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
        h1: {
            fontSize: '2.5rem',
            marginBottom: '2rem',
            color: '#4caf50',
        },
        h2: {
            fontSize: '2rem',
            marginBottom: '1.5rem',
            color: '#4caf50',
        },
    },
    components: {
        MuiButton: {
            styleOverrides: {
                root: {
                    textTransform: 'none',
                    color: '#ffffff',
                    backgroundColor: '#3e8e41',
                    '&:hover': {
                        backgroundColor: '#45a049',
                    },
                },
            },
        },
        MuiCard: {
            styleOverrides: {
                root: {
                    backgroundColor: '#28293d',
                    borderRadius: '10px',
                },
            },
        },
    },
});

export default theme;
