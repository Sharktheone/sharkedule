export type environment = {
    tags: Tag[]
    status: Status[]
    priority: Priority[]
    stage: Stage[]
    columns: Column[]
    boards: Board[]
    tasks: Task[]
    members: Member[]
    checklists: Checklist[]
    attachments: Attachment[]
    dates: Date[]
    board_names: BoardNames
    column_names: ColumnNames
    dependent_tasks: DependentTasks[]
}

export type Board = {
    name: string
    uuid: string
    columns: string[]
    tags: string[]
    description: string
    members: string[]
    priority: string
    status: string
    dueDate: string
    dates: string[]
    comments: Comment[]
    attachments: string[]
    checklists: string[]
    archived: boolean
    activity: Activity[]
    actions: string[]
}

export type Column = {
    name: string
    uuid: string
    boards: string[]
    tasks: string[]
    tags: string[]
    description: string
}

export type Task = {
    name: string
    uuid: string
    boards: string[]
    columns: string[]
    tags: string[]
    dependencies: string[]
    dependents: string[]
    comments: string[]
    description: string
    members: string[]
    priority: string
    status: string
    due_date: string
    dates: string[]
    attachments: string[]
    checklists: string[]
    done: boolean
    activity: Activity[]
}

export type Tag = {
    name: string
    uuid: string
    color: string
    icon: string
    type: string
    description: string
}

export type Member = {
    username: string
    uuid: string
    profile_picture: string
}

type Checklist = {
    name: string
    uuid: string
    items: ChecklistItem[]
    description: string
}

type ChecklistItem = {
    name: string
    uuid: string
    checked: boolean
}

type Priority = {
    name: string
    uuid: string
    color: string
    description: string
}

type Status = {
    name: string
    uuid: string
    color: string
    description: string
}

type Attachment = {
    uuid: string
    user: string
    size: number
    type: string
    date: string
    description: string
    name: string
}

export type Date = {
    uuid: string
    name: string
    date: number
    description: string
}

export type Stage = {
    name: string
    uuid: string
    color: string
    description: string
}

export type Activity = {
    //TODO
}

export type BoardNames = {
    [key: string]: string
}

export type ColumnNames = {
    [key: string]: {
        [key: string]: string
    }
}

export type DependentTasks = {
    [key: string]: {
        [key: string]: string[]
    }
}

export type NameList = {
    name: string
    uuid: string
}