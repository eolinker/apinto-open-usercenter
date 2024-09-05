export const navigateWithQueryParams = (url: string) => {
    const [path, paramsString] = url.split('?');
    const queryParams = paramsString ? parseQueryString(paramsString) : {};
    return {path, queryParams}
  }

  const parseQueryString:(queryString: string)=>{ [key: string]: string } = (queryString: string)=> {
    return queryString.split('&').reduce((acc, param) => {
      const [key, value] = param.split('=');
      acc[key] = decodeURIComponent(value);
      return acc;
    }, {} as { [key: string]: string });
  }