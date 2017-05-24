import * as React from 'react';
import * as ReactDOM from 'react-dom';
import {Provider} from 'react-redux';
import {Router} from 'react-router';
import routes from './routes';
import {History, store} from './store/Store';
import './style.css';
ReactDOM.render(
    <Provider store={ store }>
        <Router history={ History }>
            { routes }
        </Router>
    </Provider>,
    document.getElementById('app')
);
