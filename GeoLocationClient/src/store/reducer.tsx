import {List, Map} from 'immutable';
import {INITIAL_MARKERS_BATCH} from '../Utils/Actions';

interface IActions {
    type: string;
    batch: List<List<any>>;
}

const mapReducerInitialState = Map({
    dataForMarkers: List([])
});

export const mapReducer = (state = mapReducerInitialState, action: IActions) => {
    switch (action.type) {

        case INITIAL_MARKERS_BATCH:
            return state.set('dataForMarker', action.batch);

        default:
            return state;
    }
};
