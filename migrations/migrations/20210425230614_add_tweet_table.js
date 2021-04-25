exports.up = async knex =>
    knex.schema.createTable('tweet', table => {
        table.increments('id')
        table.string('content').notNullable()
        table.string('author')
        table.timestamp('created_at').defaultTo(knex.fn.now());
        table.timestamp('updated_at')
    })

exports.down = async knex =>
    knex.schema.dropTable('tweet')