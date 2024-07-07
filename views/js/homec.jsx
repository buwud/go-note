import React, { useEffect, useState } from 'react';
import { AppBar, Toolbar, Typography, Button, Container, Grid, Card, CardContent, makeStyles } from '@mui/material';

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    appBar: {
        marginBottom: theme.spacing(3),
    },
    container: {
        paddingTop: theme.spacing(4),
        paddingBottom: theme.spacing(4),
    },
    card: {
        height: '100%',
        display: 'flex',
        flexDirection: 'column',
        boxShadow: '0 2px 8px rgba(0,0,0,0.1)',
        transition: 'transform 0.3s ease-out',
        '&:hover': {
            transform: 'translateY(-4px)',
        },
    },
    cardContent: {
        flexGrow: 1,
    },
}));

const Home = () => {
    const classes = useStyles();
    const [notes, setNotes] = useState([]);

    useEffect(() => {
        fetchNotes();
    }, []);

    const fetchNotes = () => {
        fetch("http://localhost:8000/notes")
            .then(response => response.json())
            .then(data => {
                if (data.success) {
                    setNotes(data.data.rows);
                } else {
                    console.error('Error fetching notes:', data.error);
                }
            })
            .catch(error => console.error('Error fetching notes:', error));
    };

    return (
        <div className={classes.root}>
            <AppBar position="static" className={classes.appBar}>
                <Toolbar>
                    <Typography variant="h6" style={{ flexGrow: 1 }}>

                        AAAAAAAAAAAAAA
                    </Typography>
                    <Button color="inherit" onClick={handleSignIn}>Sign In</Button>
                </Toolbar>
            </AppBar>

            <Container maxWidth="lg" className={classes.container}>
                <Typography variant="h2" gutterBottom align="center">
                    FFSFDGSGDFGF to GoNote
                </Typography>
                <Typography variant="subtitle1" align="center" gutterBottom>
                    A load of notes :3
                </Typography>
                <Typography variant="body1" align="center" paragraph>
                    Sign in to take notes
                </Typography>

                <Grid container spacing={3}>
                    {notes.map((note, index) => (
                        <Grid item key={index} xs={12} sm={6} md={4}>
                            <Card className={classes.card}>
                                <CardContent className={classes.cardContent}>
                                    <Typography variant="h5" component="h2" gutterBottom>
                                        {note.title}
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary" component="p">
                                        {note.description}
                                    </Typography>
                                    <Typography variant="caption" color="textSecondary">
                                        Created at: {note.created_at}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>
                    ))}
                </Grid>
            </Container>
        </div>
    );
};

const handleSignIn = () => {
    // Handle sign-in action
};

export default Home;
