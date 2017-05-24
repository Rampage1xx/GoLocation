import * as React from 'react';
import {GoogleMap, Marker, withGoogleMap} from 'react-google-maps';
import {InfoWindowComponent} from '../InfoWindow/InfoWindowComponent';

interface IProps {
    shouldRenderInfoWindow: boolean;
}

const GettingStartedGoogleMap = withGoogleMap((props: IProps) => {
    const {shouldRenderInfoWindow} = props;

    const location = {lat: -25.363882, lng: 131.044922};

    return (
        <GoogleMap
            defaultZoom={ 3 }
            defaultCenter={ {lat: -25.363882, lng: 131.044922} }
        >
            <Marker
                position={ {lat: -25.363882, lng: 131.044922} }

                clickable={true}
            />
            <InfoWindowComponent shouldRender={ shouldRenderInfoWindow } location={ location }/>
        </GoogleMap>
    );
});

export {GettingStartedGoogleMap};
