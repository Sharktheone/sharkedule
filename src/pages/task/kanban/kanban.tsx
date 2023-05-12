import { createStyles, Text, rem } from '@mantine/core';
import { useListState } from '@mantine/hooks';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';

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

interface DndListProps {
    data: {
        position: number;
        mass: number;
        symbol: string;
        name: string;
    }[];
}

export default function Kanban() {

    const data = [
        { position: 1, mass: 1.0079, symbol: 'H', name: 'Hydrogen' },
        { position: 2, mass: 4.0026, symbol: 'He', name: 'Helium' },
        { position: 3, mass: 6.941, symbol: 'Li', name: 'Lithium' },
        { position: 4, mass: 9.0122, symbol: 'Be', name: 'Beryllium' },
        { position: 5, mass: 10.811, symbol: 'B', name: 'Boron' },
        { position: 6, mass: 12.0107, symbol: 'C', name: 'Carbon' },
        { position: 7, mass: 14.0067, symbol: 'N', name: 'Nitrogen' },
    ]
    const { classes, cx } = useStyles();
    const [state, handlers] = useListState(data);

    const items = state.map((item, index) => (
        <Draggable key={item.symbol} index={index} draggableId={item.symbol}>
            {(provided, snapshot) => (
                <div
                    className={cx(classes.item, { [classes.itemDragging]: snapshot.isDragging })}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}
                    ref={provided.innerRef}
                >
                    <Text className={classes.symbol}>{item.symbol}</Text>
                    <div>
                        <Text>{item.name}</Text>
                        <Text color="dimmed" size="sm">
                            Position: {item.position} • Mass: {item.mass}
                        </Text>
                    </div>
                </div>
            )}
        </Draggable>
    ));

    return (
        <DragDropContext
            onDragEnd={({ destination, source }) =>
                handlers.reorder({ from: source.index, to: destination?.index || 0 })
            }
        >
            <Droppable droppableId="dnd-list" direction="vertical">
                {(provided) => (
                    <div {...provided.droppableProps} ref={provided.innerRef}>
                        {items}
                        {provided.placeholder}
                    </div>
                )}
            </Droppable>
        </DragDropContext>
    );
}