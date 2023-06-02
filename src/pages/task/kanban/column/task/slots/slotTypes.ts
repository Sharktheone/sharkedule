import {
    kanbanCheckListType,
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
    | ImagesSlot
    | SubtasksSlot
    | CustomFieldsSlot
    | ChecklistsSlot



export type IndexedSlot = {
    tags: TagsSlot
    priority: PrioritySlot
    status: StatusSlot
    date_due: DateDueSlot
    stage: StageSlot
    members: MembersSlot
    progress: ProgressSlot
    images: ImagesSlot
    subtasks: SubtasksSlot
    custom_fields: CustomFieldsSlot
    checklist: ChecklistsSlot
}

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

type ImagesSlot = {
    type: "images"
    images: kanbanImageType[]
}

type SubtasksSlot = {
    type: "subtasks"
    subtasks: number //TODO: create kanbanSubtasksType
}

type CustomFieldsSlot = {
    type: "custom_fields"
    custom_fields: string //TODO: create kanbanCustomFieldType
}

type ChecklistsSlot = {
    type: "checklists"
    checklist: kanbanCheckListType[]
}