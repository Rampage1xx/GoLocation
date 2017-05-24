import * as React from 'react';
import {connect} from 'react-redux';
import {Map} from '../Components/Map/Map';
import {getMarkers} from '../Utils/HttpRequests';
export class App extends React.Component<any, any> {
    constructor(props) {
        super(props);
        const prova = getMarkers().then(r => console.log(r));
    }

    private getData() {

        const prova = getMarkers();

    }

    public render() {

        return (
            <div>
                <Map/>
            </div>
        );

    }

}

const MapStateToProps = () => undefined;
const MapDispatchToProps = () => undefined;

export const AppConnected: any = connect(MapStateToProps, MapDispatchToProps)(App);
