export type Slot = {
    type: SlotTypes
    value: string | number | string[] | undefined[]
}


export type IndexedSlot = {
    tags: string[]
    priority: string
    status: string
    date_due: string
    stage: string
    members: string[]
    progress: number
    subtasks: string[]
    custom_fields: undefined[] // TODO
    checklists: string[]
}

export type Configuration = {
    upper: SlotTypes[],
    lower: SlotTypes[],
    border?: SlotColors
    color?: SlotColors

}

export enum SlotTypes {
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