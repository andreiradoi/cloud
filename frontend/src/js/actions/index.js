
import fetch from 'whatwg-fetch'
import * as d3 from 'd3'
import { browserHistory } from 'react-router'
import FKApiClient from '../api/api.js'
import { getSampleData, updateDeepLinking } from '../utils'

import I from 'immutable'

import * as Types from './types'

function processQueryString(getState) {
    const queryString = {};
    const query = window.location.search.substring(1);
    const vars = query.split("&");
    for (let i = 0; i < vars.length; i++) {
        const pair = vars[i].split("=");
        if (typeof queryString[pair[0]] === "undefined") {
            queryString[pair[0]] = decodeURIComponent(pair[1]);
        } else if (typeof queryString[pair[0]] === "string") {
            const arr = [queryString[pair[0]], decodeURIComponent(pair[1])];
            queryString[pair[0]] = arr;
        } else {
            queryString[pair[0]].push(decodeURIComponent(pair[1]));
        }
    }
    const params = I.fromJS(queryString)
    const currentPage = getState().expeditions.get('currentPage')

    switch (currentPage) {
        case 'map': {
            if (params.has('date')) {
                return [{
                    type: Types.UPDATE_DATE,
                    date: params.get('date')
                }]
            } else if (params.has('type')) {
                return [{
                    type: Types.SELECT_FOCUS_TYPE,
                    focusType: 'documents',
                    focusID: params.get('type')
                }]
            } else if (params.has('view')) {
                const coords = params.get('view').split(',')
                const viewport = {
                    bearing: 0,
                    isDragging: false,
                    longitude: parseFloat(coords[0]),
                    latitude: parseFloat(coords[1]),
                    pitch: 0,
                    startBearing: null,
                    startDragLngLat: null,
                    startPitch: null,
                    zoom: parseFloat(coords[2])
                }
                return [{
                    type: Types.SELECT_FOCUS_TYPE,
                    focusType: 'manual'
                }, {
                    type: Types.SET_VIEWPORT,
                    viewport
                }]
            } else {
                return [{
                    type: Types.BLANK
                }]
            }
        }
        case 'journal': {
            return [{
                type: Types.UPDATE_DATE,
                date: parseFloat(params.get('date')),
                forceUpdate: true
            }]
        }
        default: {
            return [{}]
        }
    }
}

export function requestExpedition(expeditionID) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.REQUEST_EXPEDITION,
            id: expeditionID
        })

        const projectID = location.hostname.split('.')[0]

        // TODO: the API calls should be moved to middleware. The document query will need to be filtered and called whenever data is required by the frontend

        console.log('querying project')
        FKApiClient.getProject(projectID)
            .then(resProject => {
                console.log('server response received:', resProject)
                if (!resProject) {
                    console.log('project data empty')
                } else {
                    console.log('project data received')
                    dispatch({
                        type: Types.SET_PROJECT_NAME,
                        name: resProject.name
                    })
                }
            })

        console.log('querying expedition')

        FKApiClient.getExpeditions(projectID)
            .then(resExpeditions => {
                // const resExpeditions = {"name":"demoExpedition","slug":"demoexpedition"}
                console.log('server response received:', resExpeditions)
                if (!resExpeditions) {
                    console.log('expedition data empty')
                } else {
                    console.log('expedition data received, now querying documents')
                    const resExpeditionsMap = {}
                    resExpeditions.forEach(e => resExpeditionsMap[e.slug] = e)
                    FKApiClient.getDocuments(projectID, expeditionID)
                        .then(resDocuments => {
                            if (!resDocuments) {
                                resDocuments = []
                            }
                            // resDocuments = getSampleData()
                            console.log('server response received:', resDocuments)
                            if (!resDocuments) {
                                console.log('documents data empty')
                            } else {
                                console.log('documents data received')
                                const documentMap = {}
                                resDocuments
                                    .forEach((d, i) => {
                                        d.data.type = 'Feature';
                                        d.data.id = d.id
                                        d.data.date = d.timestamp * 1000
                                        if (!d.data.geometry)
                                            d.data.geometry = d.location
                                        documentMap[d.id] = d.data
                                    })
                                const documents = I.fromJS(documentMap)

                                if (documents.size > 0) {
                                    const startDate = documents.toList()
                                        .sort((d1, d2) => {
                                            return d1.get('date') - d2.get('date')
                                        })
                                        .get(0).get('date')
                                    const endDate = documents.toList()
                                        .sort((d1, d2) => {
                                            return d1.get('date') - d2.get('date')
                                        })
                                        .get(resDocuments.length - 1).get('date')

                                    const expeditions = I.fromJS(resExpeditionsMap)
                                        .map(e => {
                                            const id = e.get('slug')
                                            return e
                                                .set('id', id)
                                                .set('name', e.get('name'))
                                                .set('startDate', id === expeditionID ? startDate : Date.now())
                                                .set('endDate', id === expeditionID ? endDate : Date.now())
                                                .delete('slug')
                                        })

                                    dispatch([
                                        {
                                            type: Types.INITIALIZE_EXPEDITIONS,
                                            id: expeditionID,
                                            data: expeditions
                                        },
                                        {
                                            type: Types.INITIALIZE_DOCUMENTS,
                                            data: documents
                                        },
                                        ...processQueryString(getState)
                                    ])
                                    window.setInterval(() => {
                                        updateDeepLinking(browserHistory, getState)
                                    }, 2000)
                                } else {
                                    const expeditions = I.fromJS(resExpeditionsMap)
                                        .map(e => {
                                            const id = e.get('slug')
                                            return e
                                                .set('id', id)
                                                .set('name', e.get('name'))
                                                .set('startDate', Date.now())
                                                .set('endDate', Date.now())
                                                .delete('slug')
                                        })
                                    dispatch([
                                        {
                                            type: Types.SET_ZOOM,
                                            zoom: 2
                                        },
                                        {
                                            type: Types.INITIALIZE_EXPEDITIONS,
                                            id: expeditionID,
                                            data: expeditions
                                        },
                                        {
                                            type: Types.INITIALIZE_DOCUMENTS,
                                            data: documents
                                        }
                                    ])
                                }
                            }
                        })
                }
            })
            .catch(error => {
                console.error(error)
            // console.log('Project or expedition could not be found')
            // window.location.replace('https://fieldkit.org')
            })
    }
}

export function setViewport(viewport, manual) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.SET_VIEWPORT,
            viewport,
            manual
        })
    }
}

export function updateDate(date, playbackMode, forceUpdate) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.UPDATE_DATE,
            date,
            playbackMode,
            forceUpdate
        })
    }
}

export function selectPlaybackMode(mode) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.SELECT_PLAYBACK_MODE,
            mode
        })
    }
}

export function selectFocusType(focusType) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.SELECT_FOCUS_TYPE,
            focusType
        })
    }
}

export function selectZoom(zoom) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.SELECT_ZOOM,
            zoom
        })
    }
}

export function toggleSensorData() {
    return (dispatch, getState) => {
        dispatch({
            type: Types.TOGGLE_SENSOR_DATA
        })
    }
}

export function setMousePosition(x, y) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.SET_MOUSE_POSITION,
            x,
            y
        })
    }
}

export function setCurrentPage(currentPage) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.SET_CURRENT_PAGE,
            currentPage
        })
    }
}

export function openExpeditionPanel() {
    return (dispatch, getState) => {
        dispatch({
            type: Types.OPEN_EXPEDITION_PANEL
        })
    }
}

export function closeExpeditionPanel() {
    return (dispatch, getState) => {
        dispatch({
            type: Types.CLOSE_EXPEDITION_PANEL
        })
    }
}

export function openLightbox(id) {
    return (dispatch, getState) => {
        dispatch({
            type: Types.OPEN_LIGHTBOX,
            id
        })
    }
}

export function closeLightbox() {
    return (dispatch, getState) => {
        dispatch({
            type: Types.CLOSE_LIGHTBOX
        })
    }
}
