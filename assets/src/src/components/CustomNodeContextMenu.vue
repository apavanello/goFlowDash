<script lang="ts" setup>
import {getCurrentInstance } from "vue"
import ContextMenu from "@imengyu/vue3-context-menu"
import {getNodes} from "./Nodes.ts"
import {getEdges} from "./Edges.ts"
import {Status} from "./CustomTypes.ts"
import axios from "axios"

import Details from "./Details.vue";
import {type Node, type NodeMouseEvent} from "@vue-flow/core"

const details = Details

defineExpose({
  ShowNodeContextMenu
})

function ShowNodeContextMenu (e : NodeMouseEvent) {

    e.event.preventDefault();
    if ("clientX" in e.event) {
        ContextMenu.showContextMenu({
            x: e.event.clientX,
            y: e.event.clientY,
            items: [
                {
                    label: "Atualizar Node",
                    children: [
                        {
                            label: "Parado", onClick: () => {
                                updateNodeStatus(e.node, Status.STOP).then(() => refresh())
                            }
                        },
                        {
                            label: "Rodando", onClick: () => {
                                updateNodeStatus(e.node, Status.RUNNING).then(() => refresh())
                            }
                        },
                        {
                            label: "Concluido", onClick: () => {
                                updateNodeStatus(e.node, Status.DONE).then(() => refresh())
                            }
                        },
                    ]
                },
                {
                    label: "Detalhes",
                    onClick: () => {
                        console.log('Detalhes', details)
                        details.methods?.openDialog()
                    }
                }
            ]

        });
    }
}

async function updateNodeStatus(node: Node, status: Status) {
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
function refresh() {

    getNodes().then(() => {})
    getEdges().then(() => {})
    const instance = getCurrentInstance();
    instance?.proxy?.$forceUpdate();

    console.log('refresh')
}
</script>