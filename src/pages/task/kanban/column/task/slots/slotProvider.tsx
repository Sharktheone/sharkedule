import {createContext, ReactNode} from "react"
import {
    Configuration, CustomFieldsSlot, DateDueSlot,
    IndexedSlot,
    PrioritySlot, ProgressSlot,
    Slot,
    SlotColors,
    SlotNames, StageSlot, StatusSlot,
    TagsSlot
} from "@kanban/column/task/slots/slotTypes"
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
        if (task.checkList) slots.checklists.checklist = task.checkList

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

        if (config.border) {
            const b = config.border
            const s = slots[b]
            border = getSlotColor(s)
        }

        if (config.color) {
            const c = config.color
            const s = slots[c]
            color = getSlotColor(s)
        }

        return {upperSlot, lowerSlot, border, color}
    }

    function getSlotColor(slot: Slot): string | null {
        switch (slot.type) {
            case SlotColors.PRIORITY: {
                const prioritySlot = slot as PrioritySlot
                if (prioritySlot.priority.color) return prioritySlot.priority.color
                return null
            }
            case SlotColors.STATUS: {
                const statusSlot = slot as StatusSlot
                if (statusSlot.status.color) return statusSlot.status.color
                return null
            }
            case SlotColors.DATE_DUE: {
                const dateDueSlot = slot as DateDueSlot
                const dateDue = dateDueSlot.due_date.date
                const date = new Date(dateDue)
                const today = new Date()
                if (date > today) return "#00ff00" // TODO: Make configurable

                if (date < today) return "#ff0000"
                return null
            }

            case SlotColors.STAGE: {
                const stageSlot = slot as StageSlot
                if (stageSlot.stage.color) return stageSlot.stage.color
                return null
            }

            case SlotColors.PROGRESS: {
                const progressSlot = slot as ProgressSlot
                const progress = progressSlot.progress.percentage
                if (progress >= 100) return "#00ff00" // TODO: Make configurable
                if (progress > 80) return "#d5ff18"
                if (progress > 70) return "#ffff00"
                if (progress > 60) return "#ffcc00"
                if (progress > 30) return "#ff9900"
                if (progress > 20) return "#ff6600"
                if (progress > 10) return "#ff3300"
                return "#ff0000"
            }

            case SlotColors.CUSTOM_FIELDS: {
                return "#4433ff" // TODO: Make configurable which field to use
            }

            case SlotColors.CHECKLIST: {
                return "#4433ff" // TODO: Make configurable which checklist to use
            }

            default:
                return null
        }
    }

    return (
        <SlotContext.Provider value={slotify()}>
            {children}
        </SlotContext.Provider>
    )
}