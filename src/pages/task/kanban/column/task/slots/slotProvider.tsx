import {createContext, ReactNode} from "react"
import {Configuration, IndexedSlot, Slot, SlotNames, TagsSlot} from "@kanban/column/task/slots/slotTypes"
import {kanbanTaskType} from "@kanban/types"

type SlotContextType = {
    upperSlot: Slot[] | null
    lowerSlot: Slot[] | null
    border: string | null
    color: string | null
}

const SlotContext = createContext<SlotContextType | undefined>(undefined)

type Props = {
    children: ReactNode
    task: kanbanTaskType
}

const config: Configuration = {
    upper: [
        SlotNames.TAGS,
        SlotNames.PRIORITY,
        SlotNames.STATUS,
        SlotNames.STAGE,
        SlotNames.PROGRESS,
    ],
    lower: [
        SlotNames.DATE_DUE,
        SlotNames.MEMBERS,
        SlotNames.IMAGES,
        SlotNames.SUBTASKS,
        SlotNames.CUSTOM_FIELDS,
        SlotNames.CHECKLIST,
    ]
}


function SlotProvider({children, task}: Props) {
    function slotify() {
        let upperSlot: Slot[] = []
        let lowerSlot: Slot[] = []
        let border: string | null = null
        let color: string | null = null
        let slots: IndexedSlot = {} as IndexedSlot

        if (task.tags) slots.tags.tag = task.tags
        if (task.priority) slots.priority.priority = task.priority
        if (task.status) slots.status.status = task.status
        if (task.dueDate) slots.date_due.due_date = task.dueDate
        if (task.stage) slots.stage.stage = task.stage
        if (task.members) slots.members.members = task.members
        if (task.progress) slots.progress.progress = task.progress
        if (task.images) slots.images.images = task.images
        if (task.subtasks) slots.subtasks.subtasks = task.subtasks
        if (task.customFields) slots.custom_fields.custom_fields = task.customFields
        if (task.checkList) slots.checklist.checklist = task.checkList

        for (let slot in config.lower) {
            const slotName = config.lower[slot]
            const s: Slot = slots[slotName]
            if (s) lowerSlot.push(s)
        }

        for (let slot in config.upper) {
            const slotName = config.upper[slot]
            const s: Slot = slots[slotName]
            if (s) upperSlot.push(s)
        }
        return
    }

    return (
        <SlotContext.Provider value={{upperSlot: null, lowerSlot: null, border: null, color: null}}>
            {children}
        </SlotContext.Provider>
    )
}