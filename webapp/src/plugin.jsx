import React from 'react';

import {FormattedMessage} from 'react-intl';

import {id as pluginId} from './manifest';

import RHSView from './components/right_hand_sidebar';

import {
    ChannelHeaderButtonIcon,
} from './components/icons';

export default class Plugin {
    initialize(registry, store) {
        const {toggleRHSPlugin} = registry.registerRightHandSidebarComponent(
            RHSView,
            <FormattedMessage
                id='plugin.name'
                defaultMessage='Linkbot plugin'
            />);

        registry.registerChannelHeaderButtonAction(
            <ChannelHeaderButtonIcon/>,
            () => store.dispatch(toggleRHSPlugin),
            <FormattedMessage
                id='plugin.name'
                defaultMessage='Linkbot plugin'
            />,

        );
    }

    uninitialize() {
    //eslint-disable-next-line no-console
        console.log(pluginId + '::uninitialize()');
    }
}