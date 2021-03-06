kd = require 'kd'
KDView = kd.View
dateFormat = require 'dateformat'


module.exports = class CommentTimeView extends KDView

  constructor: (options = {}, data) ->

    options.tagName     = 'time'
    options.timeFormat ?= 'h:MM'
    options.cssClass    = kd.utils.curry 'comment-time', options.cssClass

    super options, data

  viewAppended: ->
    data = @getData()
    time = @getTime data

    @setPartial time

    @setTooltip
      title     : @getTooltipTitle data
      placement : 'above'

  getTooltipTitle: (data) ->
    relativeDates = ["Today", "Yesterday"]
    today         = new Date
    givenDate     = new Date data

    dateDifference = today.getDate() - givenDate.getDate()
    dateString     = relativeDates[dateDifference] or dateFormat givenDate, "dddd, mmmm d"
    dateString     = "#{dateString} at #{dateFormat givenDate, 'isoTime'}"

  getTime: (date) -> dateFormat date, @getOption 'timeFormat'
