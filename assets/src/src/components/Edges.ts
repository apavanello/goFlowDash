import {ref} from "vue";
import {Connection, type Edge, MarkerType} from "@vue-flow/core";
import axios from "axios";
import ShortUniqueId from "short-unique-id";

export const edges = ref<Edge[]>([]);

export const renderEdges = ref<Boolean>(true);

export async function getEdges() {
    const response = await axios.get('api/edge/list');

    edges.value.length = 0
    for (const edge of response.data) {
        console.log(edge)
        edges.value.push(updateEdgeDefaults(edge));
    }
}

function updateEdgeDefaults(edge: Edge): Edge {
    edge.labelStyle = { fill: '#10b981', fontWeight: 700 }
    edge.style = { strokeWidth: 3}
    edge.deletable = false
    return edge;
}

export function createEdge(params: Connection): Edge{
    const uid = new ShortUniqueId({length: 10});

    const edge: Edge = {
        id: uid.rnd() as string,
        type: "default",
        source: params.source,
        target: params.target,
        style: { strokeWidth: 3 },
        sourceHandle: params.sourceHandle,
        targetHandle: params.targetHandle,
        data: {},
        label: '',
        markerEnd: MarkerType.ArrowClosed
    }
    axios.post('api/edge/new', edge);
    return edge;
}
