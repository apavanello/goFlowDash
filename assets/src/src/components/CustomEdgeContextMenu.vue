<script lang="ts" setup>
import {getCurrentInstance, nextTick, ref} from "vue"
import {ContextMenu, ContextMenuGroup, ContextMenuItem} from "@imengyu/vue3-context-menu"
import {getNodes} from "./Nodes.ts"
import {getEdges, renderEdges} from "./Edges.ts"
import {EdgeFlow} from "./CustomTypes.ts"
import axios from "axios"

import {type Edge, type EdgeMouseEvent} from "@vue-flow/core"

const show = ref(false);
const edge = ref({} as Edge);


const options = ref({
  zIndex: 3,
  minWidth: 230,
  x: 500,
  y: 200
});

defineExpose({
  ShowEdgeContextMenu
})

function ShowEdgeContextMenu (e : EdgeMouseEvent) {

    e.event.preventDefault();
  if ("x" in e.event) {
    options.value.x = e.event.x;
  }
  if ("y" in e.event) {
    options.value.y = e.event.y;
  }
    edge.value = e.edge;
    show.value = true;
    console.log(edge.value)
}

async function updateEdgeFlow(edge: Edge, flowType: EdgeFlow) {


  console.log(edge)
    await axios.post(
        '/api/edge/update',
        {
            id: edge.id,
            source: edge.source,
            target: edge.target,
            sourceHandle: edge.sourceHandle,
            targetHandle: edge.targetHandle,
        },
        {
            headers: {
                'Content-Type': 'application/json',
                'Update-Type': 'edge-flow',
                'Flow-Type': flowType
            }
        }
    );
    console.log('updateEdgeFlow')
    console.log(edge)
    refresh();
}


function invertFlow() {
    console.log('invertFlow')
    updateEdgeFlow(edge.value, EdgeFlow.SERIAL)
    show.value = false;

}
function edgeDelete() {
  console.log('edgeDelete')
  axios.post(
      '/api/edge/delete',
      {
        id: edge.value.id,
      },
      {
        headers: {
          'Content-Type': 'application/json',
          'Update-Type': 'edge-delete'
        }
      }
  );
  refresh();
  show.value = false;
}


function parallelFlow() {
    console.log('parallelFlow')
    updateEdgeFlow(edge.value, EdgeFlow.PARALLEL)
    show.value = false;
}

async function refresh() {
  const instance = getCurrentInstance();
  await getNodes()
  await getEdges()
  await nextTick();
  renderEdges.value = false
  await nextTick();
  renderEdges.value = true
  console.log('refresh')
  instance?.proxy?.$forceUpdate();
  instance?.emit('refresh')

}

</script>

<template>
  <context-menu
      v-model:show="show"
      :options="options"
  >
    <context-menu-group label="Alterar Fluxo">
      <context-menu-item label="Inverter" @click="invertFlow"/>
      <context-menu-item label="Paralelizar" @click="parallelFlow"/>
  </context-menu-group>
    <context-menu-sperator/>
    <context-menu-item label="Excluir Edge" @click="edgeDelete"/>
  </context-menu>
</template>