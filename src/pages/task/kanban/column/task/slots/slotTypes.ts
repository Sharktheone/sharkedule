import {
    kanbanDateDueType,
    kanbanImageType,
    kanbanMemberType,
    kanbanPriorityType,
    kanbanStatusType,
    kanbanTagType
} from "@kanban/types"


export type Slot =
    TagsSlot
    | PrioritySlot
    | StatusSlot
    | DateDueSlot
    | StageSlot
    | MembersSlot
    | ProgressSlot
    | ImageSlot
    | SubtasksSlot
    | CustomFieldSlot


export type TagsSlot = {
    type: "tags"
    tag: kanbanTagType[]
}

type PrioritySlot = {
    type: "priority"
    priority: kanbanPriorityType
}

type StatusSlot = {
    type: "status"
    status: kanbanStatusType
}

type DateDueSlot = {
    type: "date_due"
    due_date: kanbanDateDueType
}

type StageSlot = {
    type: "stage"
    stage: string //TODO: create kanbanStageType
}

type MembersSlot = {
    type: "members"
    members: kanbanMemberType[]
}

type ProgressSlot = {
    type: "progress"
    progress: number //TODO: create kanbanProgressType
}

type ImageSlot = {
    type: "image"
    image: kanbanImageType
}

type SubtasksSlot = {
    type: "subtasks"
    subtasks: number //TODO: create kanbanSubtasksType
}

type CustomFieldSlot = {
    type: "custom_field"
    custom_field: string //TODO: create kanbanCustomFieldType
}