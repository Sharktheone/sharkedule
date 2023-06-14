import {
    kanbanSubtaskType,
} from "@kanban/types"
import {Tag, Priority, Checklist, CustomField, Progress, Member, Stage, Status} from "@kanban/types2"


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

export type TagsSlot = Tag[]

export type PrioritySlot = Priority[]

export type StatusSlot = Status[]

export type DateDueSlot = Date[]

export type StageSlot = Stage[]

export type MembersSlot = Member[]

export type ProgressSlot = Progress[]


export type SubtasksSlot = {
    type: "subtasks"
    subtasks: kanbanSubtaskType[]
}

export type CustomFieldsSlot = CustomField[]

export type ChecklistsSlot = Checklist[]