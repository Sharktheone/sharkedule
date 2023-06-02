export type kanbanBoardType = {
    uuid: string
    name: string
    description: kanbanDescriptionType
    members: kanbanMemberType[]
    tags: kanbanTagType[]
    priority: kanbanPriorityType
    status: kanbanStatusType
    dueDate: kanbanDateDueType
    dates: kanbanDateType[]
    comments: kanbanCommentType[]
    attachments: kanbanAttachmentType[]
    checkList: kanbanCheckListType[]
    images: kanbanImageType[]
    archived: kanbanArchivedType
    activity: kanbanActivityType[]
    actions: kanbanActionType[]
    columns: kanbanColumnType[]
}

export type kanbanColumnType = {
    uuid: string
    name: string
    description: kanbanDescriptionType
    tasks: kanbanTaskType[]
}

export type kanbanTaskType = {
    uuid: string
    name: string
    type: string
    members: kanbanMemberType[]
    tags: kanbanTagType[]
    priority: kanbanPriorityType
    status: kanbanStatusType
    dueDate: kanbanDateDueType
    dates: kanbanDateType[]
    description: kanbanDescriptionType
    comments: kanbanCommentType[]
    attachments: kanbanAttachmentType[]
    checkList: kanbanCheckListType[]
    images: kanbanImageType[]
    archived: kanbanArchivedType
    activity: kanbanActivityType[]
    actions: kanbanActionType[]
    subtasks: subtask[]
}

//TODO: find better way to nest subtasks - this is very ugly
type subtask = task & {
    subtasks: subtask2[]
}

type subtask2 = task & {
    subtasks: subtask3[]
}

type subtask3 = task & {
    subtasks: subtask4[]
}

type subtask4 = task & {
    subtasks: task[]
}

type task = Omit<kanbanTaskType, 'subtasks'>

export type kanbanDescriptionType = {
    description: string
}

export type kanbanPriorityType = {
    uuid: string
    priority: string
}

export type kanbanStatusType = {
    uuid: string
}

export type kanbanDateDueType = {
    uuid: string
    date: string
}

export type kanbanArchivedType = {
    archived: boolean
    date: string
    user: string

}

export type kanbanMemberType = {
    uuid: string

}

export type kanbanTagType = {
    uuid: string
    name: string
    color: string
    icon: string
}

export type kanbanDateType = {
    uuid: string,
    name: string,
    date: string,
}

export type kanbanAttachmentType = {
    uuid: string
    name: string
    type: string
    size: string
    date: string
    user: string
}

export type kanbanImageType = {
    uuid: string
    name: string
    type: string
    size: string
    date: string
    user: string
    url: string
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