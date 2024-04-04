<script lang="ts" setup>
import {ref, getCurrentInstance} from 'vue';
import {ContextMenu, ContextMenuItem} from '@imengyu/vue3-context-menu'
import {getNodes, onAddNode} from './Nodes.ts'
import {type NewNode} from './CustomTypes.ts'
import {savePanel} from "./Panel.ts";
import {FlowExportObject} from "@vue-flow/core";
import CustomDialog from "./CustomDialog.vue";
import {getEdges} from "./Edges.ts";


const show = ref(false);
const options = ref({
  zIndex: 3,
  minWidth: 230,
  x: 500,
  y: 200
});

let panelObjects = {} as FlowExportObject

function onContextMenu(e: MouseEvent, po: FlowExportObject) {
  e.preventDefault();
  panelObjects = po
  console.log('onContextMenu', panelObjects)
  //Show component mode menu
  options.value.x = e.x;
  options.value.y = e.y;
  console.log('x', options.value.x, 'y', options.value.y)
  show.value = true;
}

defineExpose({
  onContextMenu
})

const dialog = ref(false)

function onSaveDialog(data: any) {
  const newNode = {
    label: data.label,
    description: data.description,
    repo: data.repo,
    Squad: data.squad,
    StartTime: data.startTime,
    EndTime: data.endTime,
    PlanedStartTime: data.pStartTime,
    PlanedEndTime: data.pEndTime,
  } as NewNode
  console.log(newNode)

  const ePos = {x: options.value.x, y: options.value.y} as MouseEvent

  onAddNode(ePos, newNode)
  refresh()
  dialog.value = false
}

function onCloseDialog() {
  console.log('onCloseDialog')
  dialog.value = false
}

async function onSave() {
  console.log('onSave', panelObjects)
  await savePanel(panelObjects)
}

function showDialog() {
  dialog.value = true
  console.log('showDialog', CustomDialog.dialog)
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
    <context-menu-item label="Novo Node" @click="showDialog"/>
    <context-menu-sperator/>
    <context-menu-item label="Salvar" @click="onSave"/>
  </context-menu>
  <CustomDialog
      v-model:dialog="dialog"
      @onSaveDialog="onSaveDialog($event)"
      @onCloseDialog="onCloseDialog"
  >

  </CustomDialog>
</template>
