import React, { PropTypes } from 'react'
import TimelineContainer from '../../containers/common/Timeline'
import HeaderContainer from '../../containers/Root/Header'
import MapContainer from '../../containers/Map/Map'
import ExpeditionPanelContainer from '../../containers/Root/ExpeditionPanel'
import LightboxContainer from '../../containers/common/Lightbox'

class Root extends React.Component {
    constructor(props) {
        super(props)
        this.state = {}
        this.onMouseMove = this.onMouseMove.bind(this)
        this.onMouseOut = this.onMouseOut.bind(this)
    }

    onMouseMove(e) {
        const {currentExpedition, setMousePosition} = this.props
        const currentPage = location.pathname.split('/').filter(p => !!p && p !== currentExpedition)[0] || 'map'
        const x = e.nativeEvent.clientX
        const y = e.nativeEvent.clientY
        if (currentPage === 'map') {
            setMousePosition(x, y)
        }
    }

    onMouseOut() {
        const {currentExpedition, setMousePosition} = this.props
        const currentPage = location.pathname.split('/').filter(p => !!p && p !== currentExpedition)[0] || 'map'
        if (currentPage === 'map') {
            setMousePosition(-1, -1)
        }
    }

    render() {
        const {expeditionFetching, documentsFetching, documents, currentExpedition, } = this.props

        return (
            <div className="root" onMouseMove={ this.onMouseMove } onMouseOut={ this.onMouseOut }>
                { !expeditionFetching && !documentsFetching &&
                  <div>
                      <ExpeditionPanelContainer/>
                      <MapContainer/>
                      <div className="root_content">
                          <HeaderContainer/>
                          <LightboxContainer/>
                          <TimelineContainer/>
                          { !documents || documents.size === 0 &&
                            <div className="root_no-document">
                                This expedition doesn’t have any data yet.
                            </div> }
                          { this.props.children }
                      </div>
                  </div> }
                { expeditionFetching &&
                  <div style={ { color: 'black' } }>
                      fetching expedition...
                  </div> }
                { documentsFetching &&
                  <div style={ { color: 'black' } }>
                      fetching documents...
                  </div> }
            </div>
        )
    }
}

Root.propTypes = {}

export default Root
