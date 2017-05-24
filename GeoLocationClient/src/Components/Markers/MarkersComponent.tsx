import * as React from 'react';
import {Marker} from 'react-google-maps';

interface IProps {
    location: {lat: any, lng: any};
}

export const MarkersComponent = (props: IProps) => {

    return (

        <Marker
            position={ {...location} }

        />
    );
};
