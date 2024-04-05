<script lang="ts" setup>


import {ref} from "vue";
import { vMaska } from "maska"

const options = { mask: '##:##' };


const props = defineProps({
  dialog: Boolean,
  label: null,
  description: String,
  repo: String,
  squad: String,
  startTime: String,
  endTime: String,
  planedStartTime: String,
  planedEndTime: String

})

let label = ref(props.label)
let description = ref(props.description)
let repo = ref(props.repo)
let squad = ref(props.squad)
let startTime = ref(props.startTime)
let endTime = ref(props.endTime)
let planedStartTime = ref(props.planedStartTime)
let planedEndTime = ref(props.planedEndTime)


label.value = props.label
description.value = props.description
repo.value = props.repo
squad.value = props.squad
startTime.value = props.startTime
endTime.value = props.endTime
planedStartTime.value = props.planedStartTime
planedEndTime.value = props.planedEndTime




defineEmits(['onSaveDialog', 'onCloseDialog'])
</script>

<template>

  <v-dialog
      width="800"
      v-model="props.dialog"
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
                required
                v-model="label"
            >
            </v-text-field>
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
          <v-col
              cols="12"
              sm="12"
          >
            <v-text-field
                label="Squad"
                v-model="squad"
            ></v-text-field>
          </v-col>
          <v-col
              cols="6"
              sm="6"
          >
            <v-text-field
                label="Start Time"
                v-model="startTime"
                v-maska:[options]
            >
            </v-text-field>
          </v-col>
          <v-col
              cols="6"
              sm="6"
          >
            <v-text-field
                label="End Time"
                v-model="endTime"
                v-maska:[options]
            ></v-text-field>
          </v-col>
          <v-col
              cols="6"
              sm="6"
          >
            <v-text-field
                label="Planed Start"
                v-model="planedStartTime"
                v-maska:[options]
            ></v-text-field>
          </v-col>
          <v-col
              cols="6"
              sm="6"
          >
            <v-text-field
                label="Planed End"
                v-model="planedEndTime"
                v-maska:[options]
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
            @click="$emit('onCloseDialog')"
        ></v-btn>

        <v-btn
            color="primary"
            text="Save"
            variant="tonal"
            @click="$emit('onSaveDialog', {label, description, repo, squad, startTime, endTime, planedStartTime, planedEndTime})"
        ></v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

