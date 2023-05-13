import {createStyles, Group, rem} from '@mantine/core';
import test_data from "./test_data.json"
export default function Kanban() {
    const [board, setBoard] = useState<kanbanBoardType>(test_data as unknown as kanbanBoardType)

    console.log(board)

    // const items = state.map((item, index) => (
    //     <Draggable key={item.symbol} index={index} draggableId={item.symbol}>
    //         {(provided, snapshot) => (
    //             <div
    //                 className={cx(classes.item, {[classes.itemDragging]: snapshot.isDragging})}
    //                 {...provided.draggableProps}
    //                 {...provided.dragHandleProps}
    //                 ref={provided.innerRef}
    //             >
    //                 <Text className={classes.symbol}>{item.symbol}</Text>
    //                 <div>
    //                     <Text>{item.name}</Text>
    //                     <Text color="dimmed" size="sm">
    //                         Position: {item.position} • Mass: {item.mass}
    //                     </Text>
    //                 </div>
    //             </div>
    //         )}
    //     </Draggable>
    // ));


    return (
        <Group position="center" noWrap={true}>
            {board.columns?.map((column) => (
                <Column key={column.uuid} column={column}/>
            ))}
        </Group>
)

    // return (
    //     <DragDropContext
    //         onDragEnd={({destination, source}) => {
    //             console.log(destination, source)
    //
    //             handlers.reorder({from: source.index, to: destination?.index || source.index})
    //         }
    //         }
    //     >
    //         <Flex justify={"space-between"} gap={rem(16)}>
    //             <Droppable droppableId="tasks" direction="vertical">
    //                 {(provided) => (
    //                     <div {...provided.droppableProps} ref={provided.innerRef}>
    //                         {items}
    //                         {provided.placeholder}
    //                     </div>
    //                 )}
    //             </Droppable>
    //             <Droppable droppableId="in-progress" direction="vertical">
    //                 {(provided) => (
    //                     <div {...provided.droppableProps} ref={provided.innerRef}>
    //                         {items}
    //                         {provided.placeholder}
    //                     </div>
    //                 )}
    //             </Droppable>
    //             <Droppable droppableId="finished" direction="vertical">
    //                 {(provided) => (
    //                     <div {...provided.droppableProps} ref={provided.innerRef}>
    //                         {items}
    //                         {provided.placeholder}
    //                     </div>
    //                 )}
    //             </Droppable>
    //         </Flex>
    //     </DragDropContext>
    // );
}

import Column from "./column/column";
import {useState} from "react"
import {kanbanBoardType} from "./types"

const useStyles = createStyles((theme) => ({
    item: {
        ...theme.fn.focusStyles(),
        display: 'flex',
        alignItems: 'center',
        borderRadius: theme.radius.md,
        border: `${rem(1)} solid ${
            theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        padding: `${theme.spacing.sm} ${theme.spacing.xl}`,
        backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.white,
        marginBottom: theme.spacing.sm,
    },

    itemDragging: {
        boxShadow: theme.shadows.sm,
    },

    symbol: {
        fontSize: rem(30),
        fontWeight: 700,
        width: rem(60),
    },
}));
