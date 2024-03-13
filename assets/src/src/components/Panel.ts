import axios from "axios";
import {FlowExportObject} from '@vue-flow/core'


export async function savePanel(panelObjects: FlowExportObject) {
    console.log('onSave')
    console.log('obj: ', panelObjects)
    const response = await axios.post('/api/save', {
        nodes: panelObjects.nodes,
        edges: panelObjects.edges
    })



    console.log('response', response)

}