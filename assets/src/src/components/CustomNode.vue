<script setup lang="ts">
import type { NodeProps } from '@vue-flow/core'
import {Handle, Position} from '@vue-flow/core'


// props were passed from the slot using `v-bind="customNodeProps"`
const props = defineProps<NodeProps>()

</script>

<template class="vue-flow__node-default">
    <Handle id="Right" type="source" :position="Position.Right" />
    <Handle id="Top" type="source" :position="Position.Top" />
    <Handle id="Left" type="source" :position="Position.Left" />
    <Handle id="Bottom" type="target" :position="Position.Bottom" />
    <div>{{ props.label }}</div>
</template>