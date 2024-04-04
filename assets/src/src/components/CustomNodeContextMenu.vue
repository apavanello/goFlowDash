<script lang="ts" setup>
import {getCurrentInstance, ref} from "vue"
import {ContextMenu, ContextMenuGroup, ContextMenuItem } from "@imengyu/vue3-context-menu"
import {getNodes, deleteNode, updateNodeStatus} from "./Nodes.ts"
import {getEdges} from "./Edges.ts"
import {Status} from "./CustomTypes.ts"

// import Details from "./Details.vue";
import {type Node, type NodeMouseEvent} from "@vue-flow/core"

const show = ref(false);
const node = ref({} as Node);
// const details = Details


const options = ref({
  zIndex: 3,
  minWidth: 230,
  x: 500,
  y: 200
});

defineExpose({
  ShowNodeContextMenu
})

function ShowNodeContextMenu (e : NodeMouseEvent) {

    e.event.preventDefault();
  if ("x" in e.event) {
    options.value.x = e.event.x;
  }
  if ("y" in e.event) {
    options.value.y = e.event.y;
  }
    node.value = e.node;
    show.value = true;

}

async function onDeleteNode(node: Node) {
    await deleteNode(node)
    refresh();
}

async function onUpdateNodeStatus(node: Node, status: Status) {
  await updateNodeStatus(node, status)
    refresh();
}


function refresh() {

    getNodes()
    getEdges()
    const instance = getCurrentInstance();
    instance?.proxy?.$forceUpdate();

    console.log('refresh')
}
</script>

<template>
  <context-menu
      v-model:show="show"
      :options="options"
  >
    <context-menu-group label="Alterar Status">
      <context-menu-item label="Não Iniciado" @click="onUpdateNodeStatus(node, Status.STOP)"/>
      <context-menu-item label="Em execução" @click="onUpdateNodeStatus(node, Status.RUNNING)"/>
      <context-menu-item label="Pausado" @click="onUpdateNodeStatus(node, Status.PAUSE)"/>
      <context-menu-item label="Concluído" @click="onUpdateNodeStatus(node, Status.DONE)"/>
  </context-menu-group>
    <context-menu-sperator/>
    <context-menu-item label="Excluir Node" @click="onDeleteNode(node)"/>
  </context-menu>
</template>