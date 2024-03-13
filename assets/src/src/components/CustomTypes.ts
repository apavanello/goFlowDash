import {type Node} from '@vue-flow/core'


export enum Status {
    STOP = 'stop',
    RUNNING = 'running',
    DONE = 'done',
}

type NodesExtras = {
    extras: {
        status: Status
        description: string
        repo: string
    }
}

export type FullNode = Node & NodesExtras

export type NewNode = {
    label: string,
    description: string,
    repo: string,
}
