


export type kanbanType = {
    uuid: string
    name: string
    description: string
    members: string[]
    tags: string[]
    priority: string
    status: string
    dueDate: string
    dates: string[]
    comments: kanbanCommentType[]
    attachments: string[]
    checkList: kanbanCheckListType[]
    images: string[]
    archived: boolean
    activity: kanbanActivityType[]
    actions: kanbanActionType[]
    column: kanbanColumnType[]

}

export type kanbanColumnType = {
    uuid: string
    name: string
    task: kanbanTaskType[]
}

export type kanbanTaskType = {
    uuid: string
    name: string
    type: string
    members: string[]
    tags: string[]
    priority: string
    status: string
    dueDate: string
    dates: string[]
    description: string
    comments: kanbanCommentType[]
    attachments: string[]
    checkList: kanbanCheckListType[]
    images: string[]
    archived: boolean
    activity: kanbanActivityType[]
    actions: kanbanActionType[]
}

export type kanbanCommentType = {
    uuid: string
    message: string
    date: string
    user: string
}

export type kanbanCheckListType = {
    uuid: string
    name: string
    items: kanbanCheckListItemType[]
}

export type kanbanCheckListItemType = {
    uuid: string
    name: string
    checked: boolean
}

export type kanbanActionType = {
    uuid: string
    name: string
    icon: string
    color: string
    type: string
    action: string
}

export type kanbanActivityType = {
    uuid: string
    message: string
    date: string
    user: string
}