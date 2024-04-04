import {type Node} from '@vue-flow/core'


export enum Status {
    STOP = 'stop',
    RUNNING = 'running',
    DONE = 'done',
    PAUSE = 'pause',
}

export enum EdgeFlow {
    SERIAL = 'serial',
    PARALLEL = 'parallel',
}

export type NodesExtras = {
    data: NodesData
}

export type NodesData = {
    status: Status
    description: string
    repo: string
    startTime: string,
    endTime: string,
    planedStartTime: string,
    planedEndTime: string,
    squad: string,
}

export type FullNode = Node & NodesExtras

export type NewNode = {
    label: string,
    description: string,
    repo: string,
    StartTime: string,
    EndTime: string,
    PlanedStartTime: string,
    PlanedEndTime: string,
    Squad: string,
}

