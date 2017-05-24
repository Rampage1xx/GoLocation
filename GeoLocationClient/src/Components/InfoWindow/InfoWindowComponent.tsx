import * as React from 'react';
import {GoogleMap, InfoWindow, Marker, withGoogleMap} from 'react-google-maps';

interface IProps {
    shouldRender: boolean;
    location: { lat: any, lng: any };
}

export const InfoWindowComponent = (props: IProps) => {
    const {shouldRender, location} = props;
    return shouldRender ? (
        <InfoWindow position={ {...location} }>
            <p>content</p>
        </InfoWindow>

    ) : undefined;
};
