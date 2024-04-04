<script setup lang="ts">
import type { NodeProps, Node } from '@vue-flow/core'
import {Handle, Position} from '@vue-flow/core'
import {computed, ref, getCurrentInstance} from "vue";
import {type NodesData, Status} from "./CustomTypes.ts";
import CustomDialog from "./CustomDialog.vue";
import {getNodes, updateNode} from "./Nodes.ts";
import {getEdges} from "./Edges.ts";



interface Props extends NodeProps {
  data: NodesData
}

// props were passed from the slot using `v-bind="customNodeProps"`
const props = defineProps<Props>()
const show = ref(false)

const dialog = ref(false)


const changeOnStatus = computed(() => {
  switch (props.data.status) {
    case Status.STOP:
      return 'gray'
    case Status.RUNNING:
      return 'green'
    case Status.DONE:
      return 'blue'
    case Status.PAUSE:
      return 'red'
  }
})

function showDialog() {
  dialog.value = true

}
function onCloseDialog() {
  dialog.value = false
}

function onSaveDialog(data: any) {
  const node = {
    id: props.id,
    position: props.position,
    label: data.label,
    data: {
      ...props.data,
      ...data
    }
  } as Node

  updateNode(node).then(() => {
    refresh()
  })
  dialog.value = false
}

function refresh() {

  getNodes()
  getEdges()
  const instance = getCurrentInstance();
  instance?.proxy?.$forceUpdate();

  console.log('refresh')
}

</script>

<template class="vue-flow__node-custom">
    <Handle id="Right" type="source" :position="Position.Right" />
    <Handle id="Top" type="source" :position="Position.Top" />
    <Handle id="Left" type="source" :position="Position.Left" />
    <Handle id="Bottom" type="target" :position="Position.Bottom" />
  <v-card
      class="mx-1"
      max-width=100%
  >
    <v-card-title class="text-wrap" >
      <v-col>
        {{ props.label }}
      </v-col>
    </v-card-title>
    <v-card-text style="display: flex" class="pa-1">
      <v-container class="py-0">
        <v-row no-gutters>
          <v-col>
            <v-chip
                size="small"
                :color="changeOnStatus"
                variant="flat"
            >
              {{ props.data.status }}
            </v-chip>
          </v-col>
          <v-col>
            <v-chip
                size="small"
                color="primary"
                variant="flat"
                label
                style="text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;"
            >
              <v-icon icon="mdi-account-multiple" start></v-icon>
              {{ props.data.squad }}
            </v-chip>
          </v-col>
        </v-row>
      </v-container>
    </v-card-text>
    <v-card-text class="pa-1">
      <v-container>
        <v-row>
          <v-col class="px-4 py-1">
            <v-chip class="pa-1 ma-0 w-full" label>
              Start: {{ props.data.startTime}}
            </v-chip>
          </v-col>
          <v-col class="px-4 py-1">
            <v-chip class="pa-1 ma-0 w-full" label>
              End: {{ props.data.endTime}}
            </v-chip>
          </v-col>
        </v-row>
        <v-row>
          <v-col class="px-4 py-1">
            <v-chip class="pa-1 ma-0 w-full" label>
              Planed Start: {{ props.data.planedStartTime }}
            </v-chip>
          </v-col>
          <v-col class="px-4 py-1">
            <v-chip class="pa-1 ma-0 w-full" label>
              Planed End: {{ props.data.planedEndTime }}
            </v-chip>
          </v-col>
        </v-row>
      </v-container>
    </v-card-text>
    <v-card-actions>
      <v-btn
          variant="outlined"
          @click="show = !show"
      >
        Detalhes
      </v-btn>
      <v-spacer></v-spacer>
      <v-btn
          variant="outlined"
          @click="showDialog"
      >
        Edit
      </v-btn>


    </v-card-actions>

    <v-expand-transition>
      <div v-show="show">
        <v-divider></v-divider>
        <v-card-text class="details-text">
          <div>
            Detalhes:
            {{ props.data.description }}
          </div>
          <div>
            Repo:
            <a href="{{ props.data.repo }}">
              {{ props.data.repo }}
            </a>
          </div>
        </v-card-text>
      </div>
    </v-expand-transition>
  </v-card>
  <CustomDialog
      v-model:dialog="dialog"
      :label="props.label"
      :description="props.data.description"
      :repo="props.data.repo"
      :squad="props.data.squad"
      :startTime="props.data.startTime"
      :endTime="props.data.endTime"
      :pStartTime="props.data.planedStartTime"
      :pEndTime="props.data.planedEndTime"


      @onSaveDialog="onSaveDialog"
      @onCloseDialog="onCloseDialog"

  >

  </CustomDialog>

</template>
