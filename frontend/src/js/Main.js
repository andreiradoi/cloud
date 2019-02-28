// @flow weak

import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Route, Switch } from 'react-router-dom';

import log from 'loglevel';

import Map from './pages/Map';
import Source from './pages/Source';
import Files from './pages/Files';
import About from './pages/About';

import type { ActiveExpedition  } from './types';

import '../css/main.css';

type Props = {
    activeExpedition: ActiveExpedition
};

export class Main extends Component {
    props: Props

    constructor(props: Props) {
        super(props);

        log.setLevel('trace');
    }

    render() {
        return (
            <div className="main">
                <Switch>
                    <Route exact path={'/sources/:sourceId'} component={Source} />
                    <Route exact path={'/about'} component={About} />
                    <Route exact path={'/files'} component={Files} />
                    <Route exact path={'/files/:deviceId'} render={props => {
                        const { match } = props;
                        return ( <Files deviceId={match.params.deviceId} /> );
                    }} />
                    <Route path={'/'} component={Map} />
                </Switch>
            </div>
        );
    }
}

const mapStateToProps = state => ({
    activeExpedition: state.activeExpedition
});

export default connect(mapStateToProps, {
})(Main);
