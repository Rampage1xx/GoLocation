import * as React from 'react';
import {GettingStartedGoogleMap} from '../GettingStartedGoogleMap/GettingStartedGoogleMap';

class Map extends React.Component<any, any> {

    public render() {

        return (

            <div style={{height: `500px`, width: '500px'}}>
                <GettingStartedGoogleMap
                    shouldRenderInfoWindow={true}
                    containerElement={
                        <div style={{
                            width: '100%',
                            height: '100%'
                        }}/>
                    }
                    mapElement={
                        <div style={{
                            width: '100%',
                            height: '100%'
                        }}/>
                    }
                />

            </div>
        );
    }
}

export {Map};
