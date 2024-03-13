import {ref} from "vue";
import {type Edge, MarkerType} from "@vue-flow/core";
import axios from "axios";

export const edges = ref<Edge[]>([]);

export async function getEdges() {
    const response = await axios.get('api/edge/list');

    edges.value.length = 0
    for (const edge of response.data) {
        edges.value.push(updateEdgeDefaults(edge));
    }
}

function updateEdgeDefaults(edge: Edge): Edge {
    edge.labelStyle = { fill: '#10b981', fontWeight: 700 }
    edge.style = { strokeWidth: 2 }
    edge.markerEnd = MarkerType.ArrowClosed
    return edge;
}