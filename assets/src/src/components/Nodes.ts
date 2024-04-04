import axios from "axios";
import {ref} from "vue";
import {type Node} from "@vue-flow/core";
import {type FullNode, type NewNode, Status} from "./CustomTypes.ts";
import ShortUniqueId from "short-unique-id";


export const nodes = ref<Node[]>([]);

export async function getNodes() {
    const response = await axios.get('api/node/list');
    nodes.value = []

    for (const node of response.data) {
        nodes.value.push(updateNodeDefaults(node));
    }
}

function updateNodeDefaults(node: FullNode): FullNode {

    if (node.data.status == null)
        node.data.status = Status.STOP;

    if (node.data.status == Status.STOP)
        node.class = 'node-stop'

    else if (node.data.status == Status.RUNNING)
        node.class = 'animated-gradient'

    else if (node.data.status == Status.DONE)
        node.class = 'node-done'

    else if (node.data.status == Status.PAUSE)
        node.class = 'node-pause'
    //node.draggable = false

    node.deletable = false

    return node;
}

export function onAddNode(e: MouseEvent, n: NewNode) {

    const uid = new ShortUniqueId({length: 10});

    const newNode = {
        id: uid.rnd() as string,
        label: n.label,
        position: { x: e.x, y: e.y },
        type: 'custom',
        data: n
    } as FullNode

    createNode(newNode).then((data) => {console.log(data)})
    nodes.value.push(updateNodeDefaults(newNode));
}

async function createNode(node: FullNode) {

    console.log(node)
    node.data.status = Status.STOP
    const response= await axios.post('api/node/new', node);
    return response.data;
}

export async function deleteNode(node: Node) {
    await axios.post(
        '/api/node/delete',
        {
            id: node.id
        },
        {
            headers: {
                'Content-Type': 'application/json',
                'Delete-Type': 'node'
            }
        }
    );
    console.log('deleteNode')
}

export async function updateNodeStatus(node: Node, status: Status) {
    await axios.post(
        '/api/node/update',
        {
            id: node.id,
            status: status
        },
        {
            headers: {
                'Content-Type': 'application/json',
                'Update-Type': 'status'
            }
        }
    );
    console.log('updateNodeStatus')
}

export async function updateNode(n: Node) {

    const updatedNode = {
        id: n.id,
        label: n.label,
        position: n.position,
        type: 'custom',
        data: n.data
    } as FullNode

    await axios.post(
        '/api/node/update',
        updatedNode,
        {
            headers: {
                'Content-Type': 'application/json',
                'Update-Type': 'full'
            }
        }
    );
    console.log('updateNode', updatedNode)

}