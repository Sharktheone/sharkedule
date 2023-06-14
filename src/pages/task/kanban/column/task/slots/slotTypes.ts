import {
    kanbanSubtaskType,
} from "@kanban/types"


export type Slot =
    TagsSlot
    | PrioritySlot
    | StatusSlot
    | DateDueSlot
    | StageSlot
    | MembersSlot
    | ProgressSlot
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
    subtasks: SubtasksSlot
    custom_fields: CustomFieldsSlot
    checklists: ChecklistsSlot
}

export type Configuration = {
    upper: SlotNames[],
    lower: SlotNames[],
    border?: SlotColors
    color?: SlotColors

}

export enum SlotNames {
    TAGS = "tags",
    PRIORITY = "priority",
    STATUS = "status",
    DATE_DUE = "date_due",
    STAGE = "stage",
    MEMBERS = "members",
    PROGRESS = "progress",
    SUBTASKS = "subtasks",
    CUSTOM_FIELDS = "custom_fields",
    CHECKLIST = "checklists"
}

export enum SlotColors {
    PRIORITY = "priority",
    STATUS = "status",
    DATE_DUE = "date_due",
    STAGE = "stage",
    PROGRESS = "progress",
    CUSTOM_FIELDS = "custom_fields",
    CHECKLIST = "checklists"
}

export type TagsSlot = string[]

export type PrioritySlot = string

export type StatusSlot = string

export type DateDueSlot = string

export type StageSlot = string

export type MembersSlot = string[]

export type ProgressSlot = string[]


export type SubtasksSlot = {
    type: "subtasks"
    subtasks: kanbanSubtaskType[]
}

export type CustomFieldsSlot = string[]

export type ChecklistsSlot = string[]