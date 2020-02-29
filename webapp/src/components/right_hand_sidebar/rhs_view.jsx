import React from 'react';

import {makeStyles} from '@material-ui/core/styles';

import Card from '@material-ui/core/Card';

import CardActionArea from '@material-ui/core/CardActionArea';

import CardActions from '@material-ui/core/CardActions';

import CardContent from '@material-ui/core/CardContent';

import Button from '@material-ui/core/Button';

import Typography from '@material-ui/core/Typography';

export default function RHSView() {
    const classes = useStyles();

    const links = [{message: 'message1', tags: '#tag1', link: 'link1'}, {message: 'message2', tags: '#tag2', link: 'link2'}];

    return (
        <div style={style.rhs}>
            {links.map((link) => {
                return (

                    <Card
                        className={classes.root}
                        key={link.link}
                    >
                        <CardActionArea>
                            <CardContent>
                                <Typography
                                    gutterBottom={true}
                                    variant='h5'
                                    component='h2'
                                >
                                    {link.message}
                                </Typography>
                                <Typography
                                    variant='body2'
                                    component='p'
                                >
                                    {link.tags}
                                </Typography>
                                <Typography
                                    variant='body2'
                                    color='textSecondary'
                                    component='p'
                                >
                                    {link.link}
                                </Typography>
                            </CardContent>
                        </CardActionArea>
                        <CardActions>
                            <Button
                                size='small'
                                color='primary'
                            >
                                {'Go'}
                            </Button>
                        </CardActions>
                    </Card>
                );
            })}
        </div>
    );
}

const style = {
    rhs: {
        padding: '10px',
    },
};

const maxWidthConst = 345;

const useStyles = makeStyles({
    root: {
        maxWidth: maxWidthConst,
    },
});
