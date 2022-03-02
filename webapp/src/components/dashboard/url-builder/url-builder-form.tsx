import {useEffect, useState} from 'react';
import {Divider, TextField} from '@mui/material';
import {useFormik} from "formik";
import * as Yup from "yup";


export const URLBuilderForm = () => {

    console.log('builder form');

    const [resultURL, setResultURL] = useState('');

    const validationSchema = Yup.object({
        baseUrl: Yup.string()
            .url('Must be a valid url')
            .required('Base URL is required'),
        utmSource: Yup.string().trim('The UTM source cannot include leading and trailing spaces'),
        utmMedium: Yup.string().trim('The UTM medium cannot include leading and trailing spaces'),
        utmCampaign: Yup.string().trim('The UTM campaign cannot include leading and trailing spaces'),
        utmContent: Yup.string().trim('The UTM content cannot include leading and trailing spaces'),
    })
    const formik = useFormik({
        initialValues: {
            baseUrl: '',
            utmSource: '',
            utmMedium: '',
            utmCampaign: '',
            utmContent: ''
        },
        onSubmit: () => {},
        validationSchema
    });

    useEffect(() => {
        const isValid = validationSchema.isValidSync(formik.values);
        console.log("Is valid", isValid)
        if (formik.dirty && isValid) {
            const url = new URL(formik.values.baseUrl);
            if (formik.values.utmSource.length > 0) url.searchParams.append('utm_source', formik.values.utmSource)
            if (formik.values.utmMedium.length > 0) url.searchParams.append('utm_medium', formik.values.utmMedium)
            if (formik.values.utmCampaign.length > 0) url.searchParams.append('utm_campaign', formik.values.utmCampaign)
            if (formik.values.utmContent.length > 0) url.searchParams.append('utm_content', formik.values.utmContent)
            setResultURL(url.toString())
        }
    }, [formik.values])

    return (
        <>
            <form noValidate>
                <TextField
                    error={Boolean(formik.touched.baseUrl && formik.errors.baseUrl)}
                    fullWidth
                    helperText={formik.touched.baseUrl && formik.errors.baseUrl}
                    label="Base URL"
                    margin="normal"
                    name="baseUrl"
                    onBlur={formik.handleBlur}
                    onChange={formik.handleChange}
                    value={formik.values.baseUrl}
                />
                <TextField
                    error={Boolean(formik.touched.utmSource && formik.errors.utmSource)}
                    fullWidth
                    helperText={formik.touched.utmSource && formik.errors.utmSource}
                    label="UTM Source"
                    margin="normal"
                    name="utmSource"
                    onBlur={formik.handleBlur}
                    onChange={formik.handleChange}
                    value={formik.values.utmSource}
                />
                <TextField
                    error={Boolean(formik.touched.utmMedium && formik.errors.utmMedium)}
                    fullWidth
                    helperText={formik.touched.utmMedium && formik.errors.utmMedium}
                    label="UTM Medium"
                    margin="normal"
                    name="utmMedium"
                    onBlur={formik.handleBlur}
                    onChange={formik.handleChange}
                    value={formik.values.utmMedium}
                />
                <TextField
                    error={Boolean(formik.touched.utmCampaign && formik.errors.utmCampaign)}
                    fullWidth
                    helperText={formik.touched.utmCampaign && formik.errors.utmCampaign}
                    label="UTM Campaign"
                    margin="normal"
                    name="utmCampaign"
                    onBlur={formik.handleBlur}
                    onChange={formik.handleChange}
                    value={formik.values.utmCampaign}
                />
                <TextField
                    error={Boolean(formik.touched.utmContent && formik.errors.utmContent)}
                    fullWidth
                    helperText={formik.touched.utmContent && formik.errors.utmContent}
                    label="UTM Content"
                    margin="normal"
                    name="utmContent"
                    onBlur={formik.handleBlur}
                    onChange={formik.handleChange}
                    value={formik.values.utmContent}
                />
            </form>
            <Divider />
            { resultURL.length > 0 ? <TextField
                fullWidth
                multiline
                inputProps={
                    { readOnly: true, }
                }
                label="Result"
                margin="normal"
                value={resultURL}
            /> : null }
        </>
    );
};

export default URLBuilderForm;
