import {
    kanbanCheckListType, kanbanCustomFieldType,
    kanbanDateDueType,
    kanbanImageType,
    kanbanMemberType,
    kanbanPriorityType,
    kanbanStatusType, kanbanSubtaskType,
    kanbanTagType, kanbanProgressType, kanbanStageType
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
    IMAGES = "images",
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

export type TagsSlot = {
    type: "tags"
    tag: kanbanTagType[]
}

export type PrioritySlot = {
    type: "priority"
    priority: kanbanPriorityType
}

export type StatusSlot = {
    type: "status"
    status: kanbanStatusType
}

export type DateDueSlot = {
    type: "date_due"
    due_date: kanbanDateDueType
}

export type StageSlot = {
    type: "stage"
    stage: kanbanStageType
}

export type MembersSlot = {
    type: "members"
    members: kanbanMemberType[]
}

export type ProgressSlot = {
    type: "progress"
    progress: kanbanProgressType
}

export type ImagesSlot = {
    type: "images"
    images: kanbanImageType[]
}

export type SubtasksSlot = {
    type: "subtasks"
    subtasks: kanbanSubtaskType[]
}

export type CustomFieldsSlot = {
    type: "custom_fields"
    custom_fields: kanbanCustomFieldType[]
}

export type ChecklistsSlot = {
    type: "checklists"
    checklist: kanbanCheckListType[]
}