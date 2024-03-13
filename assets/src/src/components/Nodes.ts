import axios from "axios";
import {ref} from "vue";
import {type Node} from "@vue-flow/core";
import {type FullNode, type NewNode, Status} from "./CustomTypes.ts";


export const nodes = ref<Node[]>([]);

export async function getNodes() {
    const response = await axios.get('api/node/list');
    nodes.value = []

    for (const node of response.data) {
        nodes.value.push(updateNodeDefaults(node));
    }
}

function updateNodeDefaults(node: FullNode): FullNode {

    if (node.extras.status == null)
        node.extras.status = Status.STOP;

    if (node.extras.status == Status.STOP)
        node.class = 'node-stop'

    else if (node.extras.status == Status.RUNNING)
        node.class = 'animated-gradient'

    else if (node.extras.status == Status.DONE)
        node.class = 'node-done'

    //node.draggable = false
    return node;
}

export function onAddNode(e: MouseEvent, n: NewNode) {

    const id : number = nodes.value.length + 1

    const newNode = {
        id: `${id}`,
        label: n.label,
        position: { x: e.x, y: e.y },
        type: 'custom',
        extras: {
            status: Status.STOP,
            description: n.description,
            repo: n.repo,
        }
    } as FullNode

    createNode(newNode).then((data) => {console.log(data)})
    nodes.value.push(updateNodeDefaults(newNode));
}

async function createNode(node: FullNode) {

    console.log(node)
    const response= await axios.post('api/node/new', node);
    return response.data;
}

