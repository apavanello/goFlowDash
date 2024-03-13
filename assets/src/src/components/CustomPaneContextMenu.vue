<script lang="ts" setup>
import {ref} from 'vue';
import {ContextMenu, ContextMenuItem} from '@imengyu/vue3-context-menu'
import { onAddNode } from './Nodes.ts'
import {type NewNode} from './CustomTypes.ts'
import {savePanel} from "./Panel.ts";
import {FlowExportObject} from "@vue-flow/core";

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
const label = ref('')
const description = ref('')
const repo = ref('')


function onSaveDialog() {
  const newNode = {
    label: label.value,
    description: description.value,
    repo: repo.value
  } as NewNode
  const ePos = {x: options.value.x, y: options.value.y} as MouseEvent
  console.log('x', ePos.x, 'y', ePos.y)
  onAddNode(ePos, newNode)
  dialog.value = false
}

async function onSave() {
  console.log('onSave', panelObjects)
  await savePanel(panelObjects)
}

</script>


<template>
  <context-menu
      v-model:show="show"
      :options="options"
  >
    <context-menu-item label="Novo Node" @click="dialog = true"/>
    <context-menu-sperator/>
    <context-menu-item label="Salvar" @click="onSave"/>
  </context-menu>


  <v-dialog
      v-model="dialog"
      width="800"
  >
    <v-card
        prepend-icon="mdi-account"
        title="User Profile"
    >
      <v-card-text>
        <v-row dense>
          <v-col
              cols="12"
              sm="12"
          >
            <v-text-field
                label="Node Label*"
                v-model="label"
                required
            ></v-text-field>
          </v-col>
          <v-col
              cols="12"
              sm="12"
          >
            <v-text-field
                label="Descrição"
                v-model="description"
            ></v-text-field>
          </v-col>

          <v-col
              cols="12"
              sm="12"
          >
            <v-text-field
                label="Repositorio"
                v-model="repo"
            ></v-text-field>
          </v-col>
        </v-row>

        <small class="text-caption text-medium-emphasis">*indicates required field</small>
      </v-card-text>

      <v-divider></v-divider>

      <v-card-actions>
        <v-spacer></v-spacer>

        <v-btn
            text="Close"
            variant="plain"
            @click="dialog = false"
        ></v-btn>

        <v-btn
            color="primary"
            text="Save"
            variant="tonal"
            @click="onSaveDialog"
        ></v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

</template>
