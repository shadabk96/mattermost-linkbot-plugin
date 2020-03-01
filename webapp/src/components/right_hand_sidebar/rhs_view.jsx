import React from 'react';

import {makeStyles} from '@material-ui/core/styles';

import Card from '@material-ui/core/Card';

import CardActionArea from '@material-ui/core/CardActionArea';

import CardActions from '@material-ui/core/CardActions';

import CardContent from '@material-ui/core/CardContent';

import Button from '@material-ui/core/Button';

import Typography from '@material-ui/core/Typography';

import axios from 'axios';

import PropTypes from 'prop-types';

const CustomCard = (props) => {
    const maxWidthConst = 345;
    const useStyles = makeStyles({
        root: {
            maxWidth: maxWidthConst,
        },
    });
    const {data} = props;
    const classes = useStyles();
    const links = data;
    return (
        <div>
            {links.length === 0 ? (<div>{'Loading...'}</div>) : (links.map((link) => {
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
            }))}
        </div>
    );
};

CustomCard.propTypes = {
    data: PropTypes.object,
};

export default class RHSView extends React.PureComponent {
    constructor(props) {
        super(props);
        this.state = {
            links: [],
            loading: true,
        };
    }
    componentDidMount() {
        this.renderPosts();
    }
    renderPosts = async () => {
        // enter server url here
        const serverUrl = 'http://localhost:8065/';
        await axios(serverUrl + 'plugins/com.github.shadabk96.mattermost-linkbot-plugin').then((res) => {
            this.setState({links: res.data, loading: false});
        });
    }
    render() {
        if (this.state.loading) {
            return 'Loading';
        }
        return (
            <CustomCard
                data={this.state.links}
            />
        );
    }
}
