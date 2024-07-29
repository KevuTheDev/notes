-- Create a `notes` table.
CREATE TABLE notes (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(128) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_notes_created ON notes(created);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO notes (title, content, created, updated) VALUES (
    'Banana Bread Recipe',
    '2 cups (250g) all-purpose flour (spooned & leveled)\n1 teaspoon baking soda\n1/4 teaspoon salt\n1/2 teaspoon ground cinnamon\n1/2 cup (8 Tbsp; 113g) unsalted butter, softened to room temperature\n3/4 cup (150g) packed light or dark brown sugar\n2 large eggs, at room temperature\n1/3 cup (80g) plain yogurt or sour cream, at room temperature\n2 cups (460g) mashed bananas (about 4 large ripe bananas)\n1 teaspoon pure vanilla extract\noptional: 3/4 cup (100g) chopped pecans or walnuts\n',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP()
);

INSERT INTO notes (title, content, created, updated) VALUES (
    'Lorem Ipsum',
    'Lorem ipsum odor amet, consectetuer adipiscing elit. Hac ipsum ultrices augue nullam nunc. Consequat vel taciti consequat iaculis congue neque magnis placerat. Odio non elit maximus cras et. Justo id varius dignissim ultrices torquent eu tempus pharetra. Vivamus orci non dictum erat ridiculus ipsum. Eu nulla lobortis ultrices dolor parturient iaculis',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP()
);

INSERT INTO notes (title, content, created, updated) VALUES (
    'To be, or not to be - Hamlet',
    "To be, or not to be, that is the question:\nWhether 'tis nobler in the mind to suffer\nThe slings and arrows of outrageous fortune,\nOr to take arms against a sea of troubles\nAnd by opposing end them. To die—to sleep,\nNo more; and by a sleep to say we end\nThe heart-ache and the thousand natural shocks\nThat flesh is heir to: 'tis a consummation\nDevoutly to be wish'd. To die, to sleep;\nTo sleep, perchance to dream—ay, there's the rub:\nFor in that sleep of death what dreams may come,\nWhen we have shuffled off this mortal coil,\nMust give us pause—there's the respect\n",
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP()
);