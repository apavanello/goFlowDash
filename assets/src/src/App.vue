<script setup lang="ts">

import { onBeforeMount,onMounted, ref} from 'vue'
import {Background} from '@vue-flow/background'
import { ControlButton, Controls } from '@vue-flow/controls'
import {MiniMap} from '@vue-flow/minimap'
import {useVueFlow, VueFlow, type FlowExportObject, type Edge} from '@vue-flow/core'
import {nodes, getNodes} from './components/Nodes.ts'
import {createEdge, edges, getEdges, renderEdges} from './components/Edges.ts'
import axios from "axios";
import Icon from './components/Icon.vue'
import CustomNode from './components/CustomNode.vue'
import CustomNodeContextMenu from './components/CustomNodeContextMenu.vue'
import CustomPaneContextMenu from "./components/CustomPaneContextMenu.vue";
import CustomEdgeContextMenu from "./components/CustomEdgeContextMenu.vue";
import { savePanel } from './components/Panel.ts'


const { onConnect, addEdges,onNodeContextMenu, onPaneContextMenu, onEdgeContextMenu, toObject} = useVueFlow()

const dark = ref(false)
const customPaneContextMenu = ref(CustomPaneContextMenu)
const customEdgeContextMenu = ref(CustomEdgeContextMenu)
const customNodeContextMenu = ref(CustomNodeContextMenu)
let panelObjects = {} as FlowExportObject

Controls.props.position.default = "top-right"


if (import.meta.env.DEV) {
  axios.defaults.baseURL = 'http://localhost:8080'
}

onBeforeMount(() => {
  getNodes();
});

onMounted(() => {
  getEdges();
});

onNodeContextMenu((node) => {
  customNodeContextMenu.value.ShowNodeContextMenu(node);
})

onPaneContextMenu((e) => {
  panelObjects = toObject()
  console.log(customPaneContextMenu.value)
  customPaneContextMenu.value.onContextMenu(e, panelObjects)
})


onEdgeContextMenu((e) => {
  customEdgeContextMenu.value.ShowEdgeContextMenu(e)
})

onConnect((params) => {
  const newEdge: Edge = createEdge(params)
  addEdges([newEdge])
})

function logToObject() {
  console.log(toObject())
}

function toggleDarkMode() {
  dark.value = !dark.value

}

async function onSave() {
  panelObjects = toObject()
  savePanel(panelObjects)
}

</script>

<template>
  <div style="height: 100vh">
    <VueFlow
        v-model:nodes="nodes"
        v-model:edges="edges"
        :class="{ dark }"
        class="migration-flow"
        fit-view-on-init
        :default-zoom="1.5"
        :min-zoom="0.2"
        :max-zoom="4"

    >

      <Edge v-if="renderEdges"/>
      <template #node-custom="customNodeProps">
        <CustomNode v-bind="customNodeProps"  class=".vue-flow__node-default"/>
      </template>

      <Background pattern-color="#aba" :gap="8" />

      <MiniMap />

      <Controls>
        <ControlButton title="Toggle Dark Mode" @click="toggleDarkMode">
          <Icon v-if="dark" name="sun" />
          <Icon v-else name="moon" />
        </ControlButton>
        <ControlButton title="Log `toObject`" @click="logToObject">
          <Icon name="log" />
        </ControlButton>
        <ControlButton title="Save" @click="onSave">
          <Icon name="update"/>
        </ControlButton>
      </Controls>
    </VueFlow>

    <CustomPaneContextMenu ref="customPaneContextMenu"/>
    <CustomNodeContextMenu ref="customNodeContextMenu"/>
    <CustomEdgeContextMenu ref="customEdgeContextMenu"/>
  </div>
</template>
