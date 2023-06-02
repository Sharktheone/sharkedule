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
    checklist: ChecklistsSlot
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
    CHECKLIST = "checklist"
}

export enum SlotColors {
    PRIORITY = "priority",
    STATUS = "status",
    DATE_DUE = "date_due",
    STAGE = "stage",
    PROGRESS = "progress",
    CUSTOM_FIELDS = "custom_fields",
    CHECKLIST = "checklist"
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
    stage: kanbanStageType
}

type MembersSlot = {
    type: "members"
    members: kanbanMemberType[]
}

type ProgressSlot = {
    type: "progress"
    progress: kanbanProgressType
}

type ImagesSlot = {
    type: "images"
    images: kanbanImageType[]
}

type SubtasksSlot = {
    type: "subtasks"
    subtasks: kanbanSubtaskType[]
}

type CustomFieldsSlot = {
    type: "custom_fields"
    custom_fields: kanbanCustomFieldType[]
}

type ChecklistsSlot = {
    type: "checklists"
    checklist: kanbanCheckListType[]
}