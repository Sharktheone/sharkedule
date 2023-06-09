import {ReactNode, useContext, useEffect, useState} from "react"
import {Configuration, IndexedSlot, Slot, SlotTypes,} from "@kanban/column/task/slots/slotTypes"
import {Task} from "@kanban/types"
import {EnvironmentContext} from "@kanban/environment"
import {SlotContext} from "./slotContext"


type Props = {
    children: ReactNode
    task: string
}

const config: Configuration = {
    upper: [
        SlotTypes.TAGS,
        SlotTypes.PRIORITY,
        SlotTypes.STATUS,
        SlotTypes.STAGE,
        SlotTypes.PROGRESS,
    ],
    lower: [
        SlotTypes.DATE_DUE,
        SlotTypes.MEMBERS,
        SlotTypes.SUBTASKS,
        SlotTypes.CUSTOM_FIELDS,
        SlotTypes.CHECKLIST,
    ]
}


// TODO: This method of rendering tags etc is not very efficient, as it requires a lot of looping over the same data.
//  I'm a lazy b... , so I'll leave it for now, but maybe in the year 3048 or something I'll fix it - or may not KEKW.
export function SlotProvider({children, task}: Props) {
    let {environment, setEnvironment} = useContext(EnvironmentContext)

    function getTask(uuid: string) {
        return environment?.tasks?.find((task) => task.uuid === uuid)
    }

    function getPriority(priority: string) {
        return environment?.priority.find((p) => p.name === priority)
    }

    function getStatus(status: string) {
        return environment?.status.find((s) => s.name === status)
    }

    function getDate(date: string) {
        return environment?.dates.find((d) => d.name === date)
    }

    function getStage(stage: string) {
        return environment?.stage.find((s) => s.name === stage)
    }

    const [t, setT] = useState<Task | undefined>(() => getTask(task))

    useEffect(() => {
        setT(getTask(task))
    }, [environment, task])

    function slotify() {
        let upperSlot: Slot[] = []
        let lowerSlot: Slot[] = []
        let border: string | null = null
        let color: string | null = null
        let slots: IndexedSlot = {} as IndexedSlot

        if (t?.tags) {
            slots.tags = t.tags
        }
        if (t?.priority) {
            slots.priority = t.priority
        }
        if (t?.status) {
            slots.status = t.status
        }
        if (t?.due_date) {
            slots.date_due = t.due_date
        }
        if (t?.stage) {
            slots.stage = t.stage
        }
        if (t?.members) {
            slots.members = t.members
        }
        if (t?.progress) {
            slots.progress = t.progress
        }
        if (t?.subtasks) {
            slots.subtasks = t.subtasks
        }
        if (t?.custom_fields) {
            slots.custom_fields = t.custom_fields
        }
        if (t?.checklists) {
            slots.checklists = t.checklists
        }

        for (let slot in config.lower) {
            const slotName = config.lower[slot]
            const s: Slot = {
                type: slotName,
                value: slots[slotName],
            }
            if (s) lowerSlot.push(s)
        }

        for (let slot in config.upper) {
            const slotName = config.upper[slot]
            const s: Slot = {
                type: slotName,
                value: slots[slotName],
            }
            if (s) upperSlot.push(s)
        }

        if (config.border) {
            const b = config.border as string as SlotTypes
            const s = slots[b]
            border = getSlotColor(b, s)
        }

        if (config.color) {
            const c = config.color as string as SlotTypes
            const s = slots[c]
            color = getSlotColor(c, s)
        }

        return {upperSlot, lowerSlot, border, color}
    }

    function getSlotColor(type: SlotTypes, value: string | number | string[] | undefined[]): string | null {
        switch (type) {
            case SlotTypes.PRIORITY: {
                if (typeof value !== "string") return null
                const s = getPriority(value)
                return s?.color || null
            }
            case SlotTypes.STATUS: {
                if (typeof value !== "string") return null
                const s = getStatus(value)
                return s?.color || null
            }
            case SlotTypes.DATE_DUE: {
                if (typeof value !== "string") return null
                const s = getDate(value)
                if (!s) return null


                // multiply by 1000, as the unix timestamp is in seconds and JS expects milliseconds
                const date = new Date(s.timestamp * 1000)
                const today = new Date()

                if (date > today) return "#00ff00" // TODO: Make configurable

                if (date < today) return "#ff0000"
                return null
            }

            case SlotTypes.STAGE: {
                if (typeof value !== "string") return null
                const s = getStage(value)
                return s?.color || null
            }

            case SlotTypes.PROGRESS: {
                if (typeof value !== "number") return null

                const progress = value
                if (progress >= 100) return "#00ff00" // TODO: Make configurable
                if (progress > 80) return "#d5ff18"
                if (progress > 70) return "#ffff00"
                if (progress > 60) return "#ffcc00"
                if (progress > 30) return "#ff9900"
                if (progress > 20) return "#ff6600"
                if (progress > 10) return "#ff3300"
                return "#ff0000"
            }

            case SlotTypes.CUSTOM_FIELDS: {
                return "#4433ff" // TODO: Make configurable which field to use
            }

            case SlotTypes.CHECKLIST: {
                return "#4433ff" // TODO: Make configurable which checklist to use
            }
        }

        return null
    }

    return (
        <SlotContext.Provider value={slotify()}>
            {children}
        </SlotContext.Provider>
    )
}