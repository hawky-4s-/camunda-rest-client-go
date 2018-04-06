- Subscribe
- FetchAndLock
- Interceptors (BasicAuth, Request/Response logging etc -> decorator pattern)
- Topic
  - TopicName
  - Close
  - LockDuration
  - ExternalTaskHandler
  -
   /**
    * Release the topic subscription for being executed asynchronously
    *
    * @throws ExternalTaskClientException
    * <ul>
    *   <li> if topic name is null or an empty string
    *   <li> if lock duration is not greater than zero
    *   <li> if external task handler is null
    *   <li> if topic name has already been subscribed
    * </ul>
    * @return the builder


Variable handling -> JSON :)
// json.Unmarshal(textBytes, &people1)


Error Handling of Responses (how to do with type? specific ones)

Check if it's possible to restrict usage of interface{} to io.Reader for example
