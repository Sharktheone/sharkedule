import {Slot} from "@kanban/column/task/slots/slotTypes"
import {useContext} from "react"
import {SlotContext} from "@kanban/column/task/slots/slotContext"
import styles from "@kanban/column/task/slots/styles.module.scss"
import Render from "@kanban/column/task/slots/render/render"


export default function UpperSlot() {
    let context = useContext(SlotContext)
    let upperSlot = context?.upperSlot
    if (!upperSlot) return null

    let allUndefined = true
    upperSlot.forEach((slot: Slot) => {
        if (slot.value !== undefined && allUndefined) allUndefined = false
    })

    if (allUndefined) return null

    if (upperSlot.length === 0) return null


    return (
        <div className={styles.slot}>
            {upperSlot.map((slot: Slot) => (
                <Render key={slot.type} slot={slot}/>
            ))}
        </div>
    )

}